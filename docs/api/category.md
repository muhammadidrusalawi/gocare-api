# Get All Category

Used to retrieve all category data.

**URL** : `/api/categories/`

**Method** : `GET`

**Auth required** : YES - Bearer Token (JWT)

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
  "success": true,
  "message": "Categories data successfully retrieved",
  "data": [
    {
      "id": "c4a5cfa1-d6c2-4b11-8d55-234abcd9012f",
      "name": "Obat Penyakit Umum",
      "slug": "obat-penyakit-umum",
      "description": "Obat-obatan yang digunakan untuk mengobati penyakit umum seperti demam, flu, dan sakit kepala.",
      "product_count": 24
    },
    {
      "id": "a728f91e-33c2-4b74-a032-1123aabb33dd",
      "name": "Obat Khusus Anak",
      "slug": "obat-khusus-anak",
      "description": "Obat-obatan khusus untuk anak-anak, termasuk sirup dan vitamin anak.",
      "product_count": 40
    }
  ]
}
```

## Error Response

**Condition** : Unauthorized (token invalid/expired)

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Unauthorized"
}
```

# Create Category

Used to create a new category

**URL** : `/api/categories/`

**Method** : `POST`

**Auth required** : YES - Bearer Token (JWT)

**Data constraints**

```json
{
    "name": "string | required | max:255 | unique",
    "description": "string | nullable"
}
```

**Data example**

```json
{
    "name": "Obat Diabetes",
    "description": "Obat-obatan untuk mengontrol gula darah dan pengelolaan diabetes."
}
```

## Success Response

**Code** : `201 CREATED`

**Content example**

```json
{
  "success": true,
  "message": "Category created",
  "data": {
    "id": "c4a5cfa1-d6c2-4b11-8d55-234abcd9012f",
    "name": "Obat Diabetes",
    "slug": "obat-diabetes",
    "description": "Obat-obatan untuk mengontrol gula darah dan pengelolaan diabetes."
  }
}
```

## Error Response

**Condition** : Bad Request — Name invalid/empty/duplicate

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Category name already exists"
}
```

**Or**

```json
{
  "success": false,
  "message": "Category name is required"
}
```

## Error Response

**Condition** : Unauthorized (token invalid/expired)

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Token invalid"
}
```

## Error Response

**Condition** : Server Error

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```json
{
  "success": false,
  "message": "Internal Server Error"
}
```