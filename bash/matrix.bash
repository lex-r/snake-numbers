#!/usr/bin/env bash

set -eu

##
# Prints the matrix.
#
# Globals:
#   MATRIX
#   SIZE
# Arguments:
#   None
# Outputs:
#   Writes the matrix to stdout
##
matrix_print() {
    for i in `seq $(($SIZE * $SIZE))`; do
        index=$(($i - 1))
        echo -n "${MATRIX[$index]} "
        if [ $(($i % $SIZE)) -eq 0 ]; then
            echo ""
        fi
    done
}

##
# Sets the value to a field in the matrix.
#
# Globals:
#   MATRIX
#   SIZE
# Arguments:
#   Y position.
#   X position.
#   Value.
##
matrix_set() {
    index=$(($1 * $SIZE + $2))
    MATRIX[$index]=$3
}

##
# Prints debug message if debug mode is on.
#
# Globals:
#   DEBUG
# Arguments:
#   Message.
##
debug() {
    if [ "$DEBUG" == "1" ]; then
        echo $@
    fi
}

readonly DEBUG=0

declare -a MATRIX
readonly SIZE=$1
N=1

for layer in $(seq 0 $((($SIZE+1)/2-1))); do
    debug "layer $layer"
    for ptr in $(seq $layer $(($SIZE-$layer-1))); do 
        debug "→ $layer $ptr"
        matrix_set $layer $ptr $N
        N=$(($N+1))
    done
    for ptr in $(seq $(($layer+1)) $(($SIZE-$layer-1))); do 
        debug "↓ $ptr $(($SIZE-$layer-1))"
        matrix_set $ptr $(($SIZE-$layer-1)) $N
        N=$(($N+1))
    done
    for ptr in $(seq $(($SIZE-$layer-2)) -1 $(($layer))); do 
        debug "← $(($SIZE-$layer-1)) $ptr"
        matrix_set $(($SIZE-$layer-1)) $ptr $N
        N=$(($N+1))
    done
    for ptr in $(seq $(($SIZE-$layer-2)) -1 $(($layer+1))); do 
        debug "↑ $ptr $layer"
        matrix_set $ptr $layer $N
        N=$(($N+1))
    done
done

matrix_print
