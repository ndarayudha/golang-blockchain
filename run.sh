#!/bin/bash
gnome-terminal -- bash -c "make run-network; exec bash"

sleep 5

gnome-terminal -- bash -c "make run-tcp; exec bash"