# TGE Direct Buy Report

It's an application which reads Direct Buys reports (primarily constructed by `tge-spreadsheet-etl`) and creates a SWC distribution report, accepted by the `swc-distributor` tool.


Direct Buy is special use-case for handle non-automatic token subscriptions. It's for analytics, members application and `tge-direct-buy-report` tool to produce the direct-buy-report which is consumed by the `swc-distributor`  tool to distribute records.

The Direct Buy use-case is reserved for handling contributions not handled the `Pledge Queue` system.


## Data flow

Diagram: [tge-contributions-flow.pdf](https://drive.google.com/file/d/1hujmGMOZRm09xLzssf2IJtRufrifVGm0/view).

### Amounts

Please refer to the `tge-spreadsheet-etl` README file for more information and constraints about amounts.

## Status

The Direct Buy reports is created only based on **pending** records!
Please refer to the `tge-spreadsheet-etl` README file for more information.
