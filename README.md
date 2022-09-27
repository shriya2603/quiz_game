# Quiz Game

A program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.
The default time limit is 30 seconds, but should also be customizable via a flag. Quiz stops as soon as the time limit has exceeded. 
Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.
Few Assumptions:
- Questions are less than 100 
- Question and answers are not too long 

## Libraries used 
- "encoding/csv" : For reading CSV file 
- "flag" : For getting the flags data. Here we are using 2 flags : csvFile and limit . So user can provide their csv file and run the program with their desired time limit by using limit flag (in seconds )
- "strings" : For removing spaces and new line from user response
- "bufio" : For buffered I/O

## References 
- https://tutorialedge.net/golang/reading-console-input-golang/
- 
