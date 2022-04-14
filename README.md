# USER API Requirements
1.	Create a REST endpoint that allows creation of new User and returns the user id when created.
      a.	The User resource will have following fields
      i.	Full Name (Required)
      ii.	Email (required)
      iii.	Phone (optional)
      iv.	Age (optional)
2.	Create a REST endpoint that allows us to update user information
3.	Create a REST endpoint that allows us to delete a user
4.	Create a REST endpoint that allows us to search all users by email and phone and provide sorting order by field
      a.	The results should be sorted by the given field.

# To run locally:

``` 
go run main.go
```

## Open http://localhost:1323/api/v1/swagger/index.html