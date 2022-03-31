#!/bin/bash

for i in $(ls -F | grep /)
do 
    val=$(echo ${i:0:-1})
    cp ./$val/diagrams.xml diagrams_$val.xml
done

