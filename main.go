package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/subscription"
	"golang.org/x/term"
)

func getPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("password input failed: %w", err)
	}
	return strings.TrimSpace(string(bytePassword)), nil
}

func migrateSubscriptions(newPriceID string) error {
	params := &stripe.SubscriptionListParams{Status: stripe.String("active")}
	params.Filters.AddFilter("limit", "", "100")

	iter := subscription.List(params)
	for iter.Next() {
		sub := iter.Subscription()
		if len(sub.Items.Data) != 1 || sub.Items.Data[0].Price.ID == newPriceID {
			continue
		}

		_, err := subscription.Update(
			sub.ID,
			&stripe.SubscriptionParams{
				Items: []*stripe.SubscriptionItemsParams{{
					ID:    stripe.String(sub.Items.Data[0].ID),
					Price: stripe.String(newPriceID),
				}},
				ProrationBehavior: stripe.String("none"),
			},
		)
		if err != nil {
			return fmt.Errorf("failed to update subscription %s: %w", sub.ID, err)
		}
		fmt.Printf("✓ Updated subscription %s\n", sub.ID)
	}
	return nil
}

func main() {
	apiKey, err := getPassword("Enter Stripe API key (Secret Key): ")
	if err != nil {
		fmt.Printf("!! ERROR: %v\n", err)
		os.Exit(1)
	}
	stripe.Key = apiKey

	fmt.Print("Enter NEW Price ID (e.g., 'price_123abc'): ")
	var newPriceID string
	fmt.Scanln(&newPriceID)
	newPriceID = strings.TrimSpace(newPriceID)

	fmt.Print("Proceed? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
		fmt.Println("Aborted.")
		os.Exit(0)
	}

	if err := migrateSubscriptions(newPriceID); err != nil {
		fmt.Printf("!! ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Migration completed.")
}
