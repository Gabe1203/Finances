# Finances
Tracks budget and finances

## Todo/Idealog:
1. right now we will calc total based on payments, and then on shut down have the option
to manually update it -> our calcs will likely be wrong so we want this option
2. subscriptions and payments will have label, amount, billing cycle, and lasy billing date, and autoPay
-> on start up we check our payments and if autoPay is off and the billing cycle has passed check if we have paid and if it is
a different amount
3. test mode
4. implement logger to make it look pretty
5. Last logged on date -> need this so I know when to start pulling entries from 
6. payments are added to a monthly cache, at start/shut down we check the date, if the month has changed then we:
save the monthyl payments (categorized) in an excel sheet and stache it
clear the local cache of monthly paymeents and start a new cache
7. upgrade balance updates to webscraping or something similar
Connect to Bank of America API to update the value
Idea: use selenium to log into the website
8. In balances -> add a available credit part and liquidity (stock plan + checking) and then show those in cash purchasing power and liquidiity 
9. Change getAdjacentCell name

## On the source of truth for payments
For now I will use my checking account as my single source of truth. All payments considered here will be exclusively payments made by the checking account. For example, subscriptions that are paid on credit cards are just consumed by the credit card payment. Later on we can implement a feature that gives us our subscriptions

