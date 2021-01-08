# AgileEngineTest

I writed that readme after submitting the repo, I hope that its ok.

It was a fun test, because at my actual job I have been half of the time doing maintenance to a debit engine and the other half redesigning it from scratch, this background also make it difficult to not overthink about the problem.

I made a DDD or hexagonal structure project, with go-chi as router and parser.

I took the liberty to almost do it multiple accounts capable, handlers hide it from the API but the only limitation its the in-memory storage that needs a more usefull interface to save the transactions by account.

Made the account to have two different balances for credit and debit, it was not especified but I think that is the most logic approach.

When a transaction is rejected because of not enough balance it is saved rejected, I think that its also important to keep a register of them, to be able to give feedback about what failed to users.

The list of transactions finally didnt worked, I submitted like this because I surpassed the 3-4hours time, I didnt give it time to debug what happened but surely is not a big deal.

Without a doubt the most weak point are the handlers, I would like to improve the errorHandling with error types in place of just standard errors, add validations with go-playground/validator, and add logs.

Also I simplify the code by not passing the context everywhere, but a real app would need it to send metrics,logs and etc.

The project can be started doing go run main.go, or executing the build that its also in the repo.

Also the example requests can be found in the file request.http

I will be waiting your feedback.
