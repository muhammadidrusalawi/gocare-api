# Register 

Used to register a user account

**URL** : `/api/auth/register`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "name": "[name in plain text]",
    "email": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "name": "User Account",
    "email": "user@example.com",
    "password": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "The verification link has been sent to the email"
}
```

## Error Response

**Condition** : If 'email' is already registered

**Code** : `409 CONFLICT`

**Content** :

```json
{
  "success": false,
  "message": "Email already exists"
}
```

**Condition** : If verification link is sent more than once within 15 minutes

**Code** : `429 TOO MANY REQUESTS`

**Content** :

```json
{
  "success": false,
  "message": "Verification link already sent. Please wait 15 minutes"
}
```

# Verify Email

Used to verify a user's email address.

**URL** : `/api/auth/verify-email`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "verification_token": "[Verification token in plain text]"
}
```

**Data example**

```json
{
  "verification_token": "uuid-token"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "User logged in successfully",
    "data": {
      "user": {
        "id": "uuid",
        "name": "User Account",
        "email": "user@example.com",
        "role": "customer"
      },
    "token": "jwt-token"
    }
}
```

## Error Response

**Condition** : Token verification is invalid or expired

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Invalid or expired verification token"
}
```

# Login

Used to collect a Token for a registered User.

**URL** : `/api/auth/login`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "email": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "email": "user@example.com",
    "password": "secret-password"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
  "success": true,
  "message": "User logged in successfully",
  "data": {
    "user": {
      "id": "uuid",
      "name": "User Account",
      "email": "user@example.com",
      "role": "customer"
    },
    "token": "jwt-token"
  }
}
```

## Error Response

**Condition** : If 'email' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Email or password is incorrect"
}
```

# User Profile

Used to get user profile

**URL** : `/api/auth/profile`

**Method** : `GET`

**Auth required** : YES - Bearer Token (JWT)

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "User profile retrieved successfully",
    "data": {
      "id": "uuid",
      "name": "User Account",
      "email": "user@example.com",
      "role": "customer",
      "verified_at": "2026-03-08T10:12:34Z",
      "created_at": "2026-03-08T09:50:11Z",
      "updated_at": "2026-03-08T10:12:34Z"
    }
}
```

## Error Response

**Condition** : If 'token' is invalid or expired

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Unauthenticated"
}
```

# Logout

Used to invalidate a Token for a registered User.

**URL** : `/api/auth/logout`

**Method** : `POST`

**Auth required** : YES - Bearer Token (JWT)

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "User logout successfully"
}
```

## Error Response

**Condition** : If 'token' is invalid or expired

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Unauthenticated"
}
```