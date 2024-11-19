# ipusnas-bot

go mod tidy

## flow
1. user log in to the tools
2. user browse book
3. 

## flow system
1. sign in all the throwaway account, get the acces token (1 week expiration)
2. save access token to db, pair with the email, then validity status
3. pair queueing to an account

### refreshing account token (cron)
1. get all account from saved db
2. check if token account is about to expire
3. if it is, make a login request then save all the new token (? which data to save the remaining days)

### user queueing book
1. user browse or search a book
2. user place a hold on the book to tell the bot to start queueing
3. if the bot get the book, it will notify the user through email that the book is ready to be borrowed.
4. the user get the notification when the book is ready to borrow
5. the user go to the borrow screen and they asked to be prepared to go to the book menu details in the ipusnas app, and press borrow button on this bot and go back to ipusnas app to borrow the book as soon as they can (before someone else)
6. if a user failed to get the book, they must repeat to action 1
7. user can only have 1 book on place hold

### add loan queue to throwaway
1. if a user add new loan candidate
2. fetch 1 throwaway that has < 5 active loan (queueing loan or ready loan)
3. if found, add loan candidate to that throwaway

### bot borrowing book
1. fetch every throwaway that has loan book > 1.
2. loop every throaway, and loop its loan candidate book.
3. call ipusnas borrow book, if available borrow it.

## questions
1. can user add new or register throwaway account?

# NEW FEATURE SCHEME
1. this app is for queueing
2. the user can ask the app to start queueing on some books
3. then the user is notified when the book is ready to borrow, the user must open ipusnas app and this app, so they can borrow book ASAP
4. there should be a limit on how much the book can be queued per user, and the idle time when book is ready to borrow


### glosarium
1. throwaway account = account used for queueing the book in the system


###### archieved / changed / no longer used
~~3. user see the boook that containing credentials for login to the app~~
~~4. user go to the ipusnas app, sign in with the credentials~~
~~3. cron worker for update book (x)~~