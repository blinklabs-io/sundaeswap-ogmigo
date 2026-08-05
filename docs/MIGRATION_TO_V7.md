# Migrating to Ogmigo v7

The primary Ogmigo module is now github.com/SundaeSwap-finance/ogmigo/v7.
Update imports from /v6 to /v7 and use the root packages as before.

The v5 chainsync representation remains available at
github.com/SundaeSwap-finance/ogmigo/v7/ouroboros/chainsync/v5.
The v6 chainsync API remains available at
github.com/SundaeSwap-finance/ogmigo/v7/ouroboros/chainsync/v6.
The compatibility package continues to normalize v5 and v6 responses into
the primary v7 types.

Ogmios v7 adds preliminary Dijkstra and Plutus V4 support and renames the
protocol-parameter field maxReferenceScriptsSize to
maxReferenceScriptsSizePerTransaction.
