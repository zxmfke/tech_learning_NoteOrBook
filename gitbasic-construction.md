## Git Basic Construction

The basic idea of the git like the figure below

![](/assets/6.png)

---

### Preparation

Download git   
Sign up your account in github  
Prepare a directory you want to git

---

### Build
all instructions execute in git bash

##### 1.Set your github account in your local git, make the connection
    $ git config --global user.email "nono\_94@163.com" 
    $ git config --global user.name "Xiaomin Zheng"
##### 2. Go to the directory you want to git
    $ git init
   After init, this directory will be built as a repository. And you will find there is a new directory named ".git"\(invisible\).
##### 3. Add your file in directory to repository just made
    $ git add .
   use "git status" to check that all new files are added into repository
##### 4. After adding files into repository, you need to make a commitment\(every time after adding\), in order to do the version control
    $ git commit \(this is typing in vim\)
    $ git commit -m "message" \(this is typing in command line\)
   use "git log" to check commitment
   if you want to modify the commitment
    $ git commit -amend -m "message"
##### 5. All these just build in local, so this step is connecting to the github remote server.
    $ git remote add origin https://github.com/zxmfke/tech\_learning\_note.git
   The url after origin is the repository you constructed in your github. "Origin" means your remote server location.
   If you want to modify the server locaition
       $ git remote set-url origin https://github.com/zxmfke/tech\_learning\_note.git
##### 6.  So the repository is ready, we should push added file into repository\( I committed for test2.txt so the commitment is "version 2"\)
     $ git push origin HEAD
   "Origin" is defined before, "HEAD" means the latest version.
   
   
All these steps is only for build your git up, and you can update your file or code.
