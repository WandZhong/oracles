## Mechanics structure

Structure:

1. Oracle builds the subscription queue.
1. At any given point in time, there is only one current subscription queue.
1. User is not able to choose into  which tranche he pleges. Oracle always puts the user into the current subscription queue.
1. We don't have future subscription queues (we have only the current one and the past ones).
1. We are able to browse all subscription queues.

Parameters:
1. Oracle has information about maximum user contribution per tranch (`MaxUserContribution`).
1. Oracle has information about maximum total user contribution (`MaxTotalUserContribution`) - which is the maximum amount of SWC user can buy overall.
1. The parmaters can be changed.


## SWC subscription mechanism

### Subscribing:

1. User transfer his BRG* to a SWCsubscription contract
2. The oracle listen on BRG* token events
3. Once there is a transfer event to SWCsubscription it adds the user for the upcoming SWT tranche release in database
4. The oracle sends back any token amount in excess of the maximum total user contribution


### Cancellation

1. User can ask the SWCsubscription contract to cancel a certain amount of BRG he plegdege in the current subscription queue.
1. The SWCsubscription contract emits a cancelation request event
1. The swc-subscription oracle will listen on SWCsubscription events and upon such an event the oracle checks the request.
	1. If the request exceed the current contribution it will be rejected (no adjustemnt is made).
	1. Otherwise it is authorized.
1. The authorization by the oracle will trigger the SWCsubscription contract to transfer back the requested BRG tokens to the user's account.
1. After authorizing the cancellation the oracle will reduce the user's subscriptions starting with the most recent one.



### SWT tranche distribution

1. The oracle will construct the distribution list based on the subscription order: `L = [l1, l2, ..., ln]`.
The queue is capped to the new tranche size. Requests which exceeds the distribution list (when `L` covers all tranche size), will be moved to the new subscription queue for the next tranche.
1. SWCsubscription will burn the amount of BRG pledged for the newly released SWC.
