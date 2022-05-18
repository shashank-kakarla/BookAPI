
# Book API

The project is a Book API written in Golang using libraries like Gin and GORM. This API performs all the **CRUD** operations like -
1. Create a New Book with Title and Author.
2. Retrieve all the Books that exist in the Database.
3. Retrieve the Books with the given ID (automatically generated at creation).
4. Retrieve the Book with the given Title.
5. Retrieve all the Books with the given Author.
6. Delete a Book with given ID.
7. Update the Title or Author of a Book with given ID.

## Packages
#### Gin
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. For further information on the package, you can checkout - [Gin Golang](https://github.com/gin-gonic/gin)

#### GORM
GORM provides CRUD operations and can also be used for the initial migration and creation of the database schema.For further information on the package, you can checkout - [GORM Golang](http://gorm.io)

## API Reference

#### Get all items

```http
  GET /books
```
![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/GET_Books.gif)

#### Get Book with given ID

```http
  GET /books/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `uint` | **Required**. Id of item to fetch |

![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/GET_Books_ID.gif)

#### Get Book with given Title

```http
  GET /books/title/:title
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Id of item to fetch |

![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/GET_Books_Title.gif)

#### Get Book with given Author

```http
  GET /books/author/:author
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `author`      | `string` | **Required**. Id of item to fetch |

![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/GET_Books_Author.gif)

#### Create Book with given Author and Title
This operation here also prevents creation of duplicate data.

```http
  POST /books
```

* Request Body Example -
```
{
    "title": "Fairy tales",
    "author": "Hans Christian Andersen"
}
```
![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/POST_Books.gif)

#### Update Book with given Author and Title for the ID

```http
  PATCH /books/:id
```
* Example 1
```
{
    "title": "Fairy tales"
}
```
* Example 2
```
{
    "author": "Stephenie Meyer"
}
```
* Example 3
```
{
    "title": "Twilight"
    "author": "Stephenie Meyer"
}
```
![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/PATCH_Books.gif)

#### DELETE Book with given Author and Title for the ID

```http
  DELETE /books/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `uint` | **Required**. Id of item to delete |

![](https://github.com/shashank-kakarla/BookAPI/blob/master/media/DELETE_Books.gif)


## Deploy BookAPI
Developed dockerfile to build and run the bookapi service, Inorder to attach Database used MySql as a database and implemented docker-compose

### Pre-requisites
  1. docker
  2. docker-compose

**Build & Deploy Book-API**

```shell
    docker-compose build && docker-compose up -d
```
**Shut-Down the application**

```shell
docker-compose down --remove-orphans --volumes 
```

