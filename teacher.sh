#!/bin/bash

INTERVIEW_NUMBER=$(find /path/to/mystery/folder -type f -name "interview*" | sed 's/.*interview\([0-9]*\).*/\1/')
export INTERVIEW_NUMBER

echo "INTERVIEW_NUMBER=$INTERVIEW_NUMBER"


INTERVIEW_FILE="/path/to/mystery/folder/interview$INTERVIEW_NUMBER"
echo "Interview $INTERVIEW_NUMBER contains:"
cat "$INTERVIEW_FILE"
echo

echo "MAIN_SUSPECT contains: $MAIN_SUSPECT"
