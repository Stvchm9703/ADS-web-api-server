# Advanced Database System Project

## Content

- ##### [System Setup](/doc#setup)
  - [Development](/doc/#setup/develop)

- ##### [Data Structure and Api Information](/doc/#api)
  - [Department](/doc/#api/department)
  - [Course](/doc/#api/course)
  - [Offer](/doc/#api/offer)
  - [Student](/doc/#api/student)
  - [Enroll](/doc/#api/enroll)



------

## Summary 

This project is aiming for using MongoDB as NoSQL Application,

## Structure of this system 
This system program are using serval tools and language
- `Nodejs / Nuxt.js` for front-end SPA prerendered web app 
- `Golang` for back-end REST-ful API server and static resourse server for PreRender web app
- `MongoDB` for main storage NoSQL program 

### Framework that we consider 
While the Framework of this program is using `Golang` as api-server 
for data input / output guardian, to prevent the data structure of each record
may not have huge impact change, when front-end program have different version of data model.\
Compare to direct accessing to MongoDB by Web App, more high level of security checking are necessary to handle by server-side while end-user get resource 

### Using Golang as server-side program
Compare to NodeJS, Python and Golang, Golang have more stable performance.
Although Nodejs and Python are interpreted scripting languages and have higher flexibly, 
as production program, Golang have strict variables requirement during compiling that giving a good memory management, 
while it have less complexity for the file structure then Java / C#.

Plus, Golang have a better package dependecy management tool,`go mod`, in officially.
Less complex than NodeJS's `npm` and C#'s `NuGet`

### Using Nuxt.js as client-side program 
Indeed, Nuxtjs javascript framework is based on Vue.js structure to develop as NodeJS Program that focusing on Server-Side Rendering.
However, it is also a good framework for prerendering. Compare to Vue, it provides a standard file structure for web application.

### Using MongoDB as Data Storage Program
MongoDB is a document database, which are more flexibly on handle data storage. Ignore the data type and structure, so that giving more time focus on front-end developing.

---

