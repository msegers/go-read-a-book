# Go-read-a-book

> Note
> This was a simple, project I did as part as an interview, at the time of writing the code (and now honestly) I've been a beginner with Go.

A simple project which compiles an application which can read data from any posted string. Can be quite long.
By default you can post this data to localhost:5555, or you can set the port when starting up by adding a -tport flag.

e.g: go-read-a-book.exe -tport 1001

Information about the parsed text can be fetched at localhost:8080/stats this port can be configured just like the reading port by adding a -wport flag followed up by the port. Also when adding a get param ?N=10 you get the top 10 words and chars instead of 5 (of course 10 can be any int).
