# fo76-market-tool

# PROJECT IS OUTDATED
the mod received unofficial updates by [MsZelia](https://github.com/MsZelia) ([Nexus](https://next.nexusmods.com/profile/MsZelia?gameId=2590))
- [Invent O Matic Pipboy (Unofficial) Update](https://www.nexusmods.com/fallout76/mods/2324)
- [Invent O Matic Stash (Unofficial) Update](https://www.nexusmods.com/fallout76/mods/2335)

also, [MsZelia](https://github.com/MsZelia) created a [python script](https://github.com/MsZelia/InventOMatic-Parser), which convert `itemsmod.ini` to table, which can 
be transfered to [Google Spreadsheet](https://www.google.com/search?q=google+sheets), [Microsoft Excel](https://www.microsoft.com/en-us/microsoft-365/excel), etc

**TL;DR:**
- start use new mods ([1](https://www.nexusmods.com/fallout76/mods/2324), [2](https://www.nexusmods.com/fallout76/mods/2335))
- convert mod's result via [script](https://github.com/MsZelia/InventOMatic-Parser)
- import [script](https://github.com/MsZelia/InventOMatic-Parser)'s result in your spreadsheets software

## Description

Tool which help you gather and process data from [Invent O Matic Stash](https://www.nexusmods.com/fallout76/mods/698).

## Features

- transform data gathered from [Invent O Matic Stash](https://www.nexusmods.com/fallout76/mods/698) to feed it to [fo76market website](https://fo76market.herokuapp.com/#/).

## Usage/Examples

### Transform
```bash
./fo76_market_tool transform "C:\Program Files (x86)\Steam\steamapps\common\Fallout76\Data\itemsmod.ini" "C:\result.json"
```

