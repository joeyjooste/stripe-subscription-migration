## Stripe Subscription Migration Script

Welcome, this is a very simple script.

It takes a product and migrates all existing prices to a new price.

Example:

Lets say you have Product A.

Product A is being sold for $20 and $30 per month.

But you want to rather move everyone to $25 per month.

You would go into stripe and create a new Price for your product.

Now you realise you need a script to move the customers.

That's where this script comes in.

Simply enter the new Price_ID and it migrates everyone over.

## How to use

Install Go:

https://go.dev/doc/install

Once install follow these commands:

```
git clone https://github.com/joeyjooste/stripe-subscription-migration.git
```

```
cd stripe-subscription-migration
```

```
go run main.go
```

Follow the prompts and enter your API keys


## Notes:

Migrations currently only do 100 at a time, so you may have to run the script multiple times

!!IMPORTANT!!: I would highly recommened that you run this on test data first, using stripes test mode before you run it over your live customer subscription. This assures you that its acting as intended

## Roadmap:

Paginination approach which does 100 at a time so you can migrate thousands in one go.

Additional Error Handling for edge cases.

Ability to handle subscriptions other than "active" subscriptions 

