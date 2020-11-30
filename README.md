# Quote Generator 

Here I am going to try to build a simple API which will add a quote to an image. This can come in a number of different ways. 


Ultimately I'd like to add a simple web interface that will work with the rest of the code to either 
generate an image with a quote or it will pick a new image from 
Unsplash and provide the user with the necessary wonders. 

## TODO 
* Pull quotes from a CSV / [POSTGRES](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/) table to be placed into the file 
* Pull random pictures from Unsplash, resize image, and then generate 
* Create API framework to result in an image to generate the image 
* Check relational font size and image size with however long the actual image (think how powerpoint does a great job to fit text into the size)
* https://blog.golang.org/using-go-modules
* https://github.com/hbagdi/go-unsplash
* https://github.com/fogleman/gg/blob/master/examples/sine.go