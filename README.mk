# Sweetbridge Oracles

This repository contains blockchain Oracles used by Sweetbridge project.


## Scripts

### abigen

`scripts/abigen.js` is used to generate Go interfaces based on JSON interface file produced by Truffle. Whenever new contracts are provided you should regenerate Go interfaces

    node scripts/abigen.js <path to the truffle contract builds>

Usually you should just use make:

    make abigen-backsage
