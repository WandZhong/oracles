# eth-contribs-etl

Eth-Contribs_ETL is a command that extracts transactions from an ethereum blockchain and loads them into the a transactions. Only the transactions related to the filter addresses are loaded.
This program processes blocks at a rate of 1/second.
The validation block rate in mainet is about 4block/min, or an average of 240 blocks/hour, or an average of 6000 blocks a day.
That means that the program can process one day of block validation (6000) in 1h40min.

The more specific documentation:

* [TGE Ethereum Contributions ETL](https://docs.google.com/document/d/1h9-4Hm16ygKrzVdwWO9nPu8sVO6RriPgDI11nZ_RTpY/edit#heading=h.xnl5gebzkyed)

## Usage

Run command with `-h` to check the usage.

The program requires:

+ Access to the database that contains the stored function that records transactions
+ The DB User used for the connection, must have stored function execution access
+ It is possible to use an ini file that contains all the options


## TODO

+ [ ] Use a file to contain the list of addresses to filter on
+ [ ] To be more dynamique, use an dynamic way to retrieve addresses of interest
