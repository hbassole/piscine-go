#! /bin/bash
ls -tF | tr '\r\n' ',' | sed 's/\(.*\),/\1/' | sed 's/*//g'