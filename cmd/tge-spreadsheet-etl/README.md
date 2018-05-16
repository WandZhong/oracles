# TGE Spreadsheet ETL

It's an application which loads members Sweetcoin contribution and stores them as a Direct Buy records into the Postgresql database.

This is a limited usage tool to be used only for a certain set of contributions, verified and managed by the TGE & Compliance Teams for the early TGE stages.

Please refer for to ./cmd/tge-direct-buy-report/README.md for more information.


## Data flow

Diagram: [tge-contributions-flow.pdf](https://drive.google.com/file/d/16vjNoWsjhVuFHbiOhdsppq2toc82tpUA/view?usp=sharing).

## Status

We support following statuses:

+ "on hold"  -- order acknowledged but not scheduled for distribution
+ "pending"  -- scheduled for the TGE distribution
+ "done"  -- distribution is done.

The direct buy operator can update status between _on hold_ and _pending_ (in any way). But once status is set to **done** it can't be changed.

The Direct Buy records can be updated. However not fields are overwritten:

+ we keep old timestamp values
+ we don't allow to set a lower status then the one recorded in DB. The **status is always** the maximum of the status in DB and the status in the spreadsheet.
