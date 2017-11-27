# tranche-queue

The application is responsible for listening on the Blockchian `BRG Transfer`, `SWCqueue LogDirectPledge` and `SWCqueue LogCancel` events to manage the Pledge Queue.

The technical documentation is available in google docs: [Sweetcoin Queue Oracle](https://docs.google.com/document/d/1_aAedNVC8YngK_xD8qeP4_CGeBlGS2fMexu3S7ltnYk/)

The oracle is an ephemeral process and requires PostgreSQL database to build and update Tranche Release Queue.


## TODO

+ [ ] KYC check
+ [ ] fault resistance
+ [ ] check past events
+ [ ] max value control
+ [ ] withdraw
