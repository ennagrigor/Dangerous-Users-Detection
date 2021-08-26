# Dangerous User Detection on Social Media

<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/Logo.png" width=200>

### A final project a Ariel University 

The purpose of the product is to scan social media platforms that have 
public APIs and extract text via posts, comments, tweets etc. 
Then it will monitor and identify keywords according to a pre-defined "dangerous" vocabulary (that can always be altered in order to fit any spesific need). 
With the help of monitoring and identifying we can maintain a list of suspicious people who may pose a social danger.

**Created by:**

[Enna Grigor](https://github.com/ennagrigor)

## Languages

Written in GO lang for server side and REACT for client side. 

## Database

<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/blevelogo.png" width=200>
The database that was used is Bleve as it is build in with go lang and fit the project needs. 

## The Server Side

The project focused on Twitter API and extracting tweets in order to locate dangerous users. 
It feachers the folowing: 

- `TextDetection` - That feachers the dictionaries (can be at any language). 
- `SocialMedia` - Interface.
- `Scheduler` - The scheduler that runs the app in intervals.
- `Router` -That feachers the router and handlers.
- `Configurations` -That should be altered by each user. 

### The Dictionary: 

Is a number of  json files (each file is a different language) that holds our vocabulary and score. 
- `The score `  Is a number with values from 1 to 10 that tell use how dangeous the word is. 

Example of the "english dictionary":

```json
    {
    "bomb": 10,
    "bombing": 10,
    "explode": 9,
    "explosion": 9,
    "stab": 9,
    "stabing": 9,
    "stabed": 8,
    "terror": 10,
    "terror attack": 10,
    "attack": 4
    }
```
### The classification:

The classification is done by a "score" method. 
Each word in the vocabulary has a number from 1-10 and a user will have the value "threat" added every time we find a suspicious tweet according to the score of the tweet. 

- `no risk` - If the score is 0 ( user will not be saved in data base).
- `low risk` - Score 1 to 5 (not including).
- `medium risk` - Score 5 to 10 (not including). 
- `high risk` - Above 10 (including).

### Top 5 Users

The app will do a check and present the top 5 most dangerous users. 
This is done by checking all the tweets found by the userID and summing up their score. 
The five users with the hiest score will be shown to the client. 

<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/Top5.png" width=600>
 
## The client Side: 
 
 The client will be able to see the top 5 users the were classified as most dangerous: 
 
<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/Top5.png" width=1000>
 
 The will also see the 10 most recent tweets that were found including information about the user and score: 
 
<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/Tweets.png" width=1000>
 
 And they can search for tweets from the database based on the parameters from the chart: 
 
<img src="https://github.com/ennagrigor/Dangerous-Users-Detection/blob/master/Images/search.png" width=1000>
 


