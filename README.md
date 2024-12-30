# fileair ✈
For a while now I have used services like workupload and weshare. They are nice and easy to use, but at least the free tier has its limitations. 
I recently came across the task of sharing (for my standarts) large files, 60GB+ to be percise. And I didn't want to signn up for yet another monthly subscription.
So I did the only thing I had left, I've written a CLI for file sharing. Super staright forward. All with one command. 

Since I mostly live in my terminal anyway, might as well share files from there. 

Now all that is left is cloud storage, and who better to ask for some than AWS. 

So I've set up a bucket for myself and connected it to the CLI.

## Usage 
- Create a bucket + IAM user with bucket read/write rights
- Add the .env file with your credentials

### Command
The command is as follows <br>
`<pathToFile> <h|d> <number>` <br>
To elaborate on that, 
1. `<pathToFile>` represents the path to the file you want to uplaod, if you want to upload a folder, tutn it into a zip file (cheaper and faster to upload)
2. `<h|d>` is a parameter used to create a signed link which will expire after h: hour or d:day
3. `<number>` specifies the number of hours or days until the link expires

The link will atomatically be saved to your clipboard, no need to copy paste the returned link. It's just there for future refernce. 

I also have implemented a piecemeal upload for files larger than 100MB, so very large files can be uploaded as well. 

An example usage could be <br>
`fileair /Users/ninja/desktop/new.txt h 2` <br>
Wich would (after confirmation) upload the file to your AWS bucket and produce a signed link valid for 2 hours.

A word of warning, AWS will charge for usage, this service is not free, but not a monthly subscription either. You pay as you use. 
