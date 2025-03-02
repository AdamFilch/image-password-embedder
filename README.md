# image-password-embedder


# Main Idea

Just like pirate software saving passwords in images

we can have scripts that can encode and decode an image

first it will be a CLI tool, int he future can be a ui so user can open and use the gui to do encryption and choosing image


# Flow

1. User inputs an image
2. Adds an encryption key
3. Puts in a password
4. The script embeds the password into the image using the encryption key


5. User chooses an image to decode
6. Uses the same encryption key that is used to embed password
7. The scripts grabs the password from the image
8. user chooses to show the password on to the screen or just copy the password straight to your clipboard



## Other Add Ons?

1. User should be able to create their own encryption mechanism but the way the password is saved into the images might have to be the same
2. the passwords saved will be stored in a folder somwehre in the user's computer. might need to hide the folder incase of accidental deletion
3. user needs to log in to access script
4. MFA? so user can add more layers



## Raised Questions
Can you hide algorithms in code?
Like hide how my code embeds the message into the image?

i think its fine if we cannot hide the algorithm, if we use the encryption key the user provided to store the password into the image like a query ?encrypt-key=password and we store this inside the image, i think its would be hard to get the data from the image without knowing this encrypt key

then to take it out, the user just needs to enter the same encryption key so the script knows what it is searching for.