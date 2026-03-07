# Register 

Used to register a user account

**URL** : `/api/auth/register/`

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
    "message": "Kode Otp berhasil dikirim, silahkan cek email Anda",
    "data": null
}
```

## Error Response

**Condition** : If 'email' is already registered

**Code** : `409 CONFLICT`

**Content** :

```json
{
  "success": false,
  "message": "Email sudah terdaftar"
}
```

**Condition** : If 'OTP Code' is sent more than 5 times in 5 minutes

**Code** : `429 TOO MANY REQUESTS`

**Content** :

```json
{
  "success": false,
  "message": "OTP sudah dikirim. Cek email atau coba lagi dalam 5 menit"
}
```

# Verify OTP

Used to verify OTP code

**URL** : `/api/auth/verify-otp/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "email": "[valid email address]",
    "otp_code": "[OTP code in plain text]"
}
```

**Data example**

```json
{
    "email": "user@example.com",
    "otp_code": "123456"
}
```

## Success Response

**Code** : `201 CREATED`

**Content example**

```json
{
    "success": true,
    "message": "Berhasil registrasi",
    "data": {
        "id": "uuid",
        "name": "User Account",
        "email": "user@example.com",
        "role": "customer"
    }
}
```

## Error Response

**Condition** : If 'OTP code' is wrong or expired

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Kode OTP salah atau kadaluarsa"
}
```

**Condition** : If 'OTP code' is wrong more than 5 times

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Kode OTP terlalu banyak salah, coba lagi dalam 5 menit"
}
```

# Login

Used to collect a Token for a registered User.

**URL** : `/api/auth/login/`

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
    "password": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Berhasil login",
    "data": {
      "user": {
        "id": "uuid",
        "name": "User Account",
        "email": "user@example.com",
        "role": "customer"
      },
      "token": "93144b288eb1fdccbe46d6fc0f241a51766ecd3d"
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
  "message": "Email or password salah"
}
```

# User Profile

Used to get user profile

**URL** : `/api/auth/profile/`

**Method** : `GET`

**Auth required** : YES - Bearer Token (JWT)

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "Berhasil logout",
    "data": {
      "id": "uuid",
      "name": "User Account",
      "email": "user@example.com",
      "role": "customer"
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
  "message": "Token tidak valid atau sudah kadaluarsa"
}
```

# Logout

Used to invalidate a Token for a registered User.

**URL** : `/api/auth/logout/`

**Method** : `POST`

**Auth required** : YES - Bearer Token (JWT)

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "Berhasil logout",
    "data": null
}
```

## Error Response

**Condition** : If 'token' is invalid or expired

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Token tidak valid atau sudah kadaluarsa"
}
```

# Forgot Password

Used to send a reset password link to the user's email

**URL** : `/api/auth/forgot-password/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "email": "[valid email address]"
}
```

**Data example**

```json
{
    "email": "user@example.com"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "Link reset password terkirim. Silahkan cek email.",
    "data": null
}
```

## Error Response

**Condition** : If 'email' is not found or invalid

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "success": false,
  "message": "Email tidak ditemukan"
}
```

**Condition** : If 'email' is sent in less than 15 minutes

**Code** : `429 TOO MANY REQUESTS`

**Content** :

```json
{
  "success": false,
  "message": "Link reset password sudah dikirim. Cek email."
}
```

# Confirm New Password

Used to confirm new password after visit reset password link (redirect to client url with token)

**URL** : `/api/auth/confirm-new-password/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "new_password": "[new password in plain text]",
    "token": "[token in plain text]"
}
```

**Data example**

```json
{
    "new_password": "abcd1234",
    "token": "93144b288eb1fdccbe46d6fc0f241a51766ecd3d"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "success": true,
    "message": "Password baru berhasil dibuat",
    "data": null
}
```

## Error Response

**Condition** : If 'token' is invalid or expired

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "success": false,
  "message": "Token tidak valid atau sudah kadaluarsa"
}
```