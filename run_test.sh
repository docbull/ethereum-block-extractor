#!/bin/bash
# INLab, Dongguk Univ.

geth --datadir /home/inlab/inlab-ethereum-hp2p/data/ init genesis.json
geth --networkid 15 --datadir data --port 30331 --mine --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000 console 2>> /home/inlab/inlab-ethereum-hp2p/data/client1.log
