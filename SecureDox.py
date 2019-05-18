#! /usr/bin/env python3

import subprocess

# Choose Encryptionnir Decryption

answer = input("Encrypt or Decrypt? ")

# Encryption

if (answer == "Encrypt"):

  data = input("What is the path of the data file you would like to encrypt? ")

  new_file = input("What would you like to call the new encrypted file? ")

  pubkey = input("What is the path of the public key? ")


  subprocess.run(["openssl", "rand","-out", "randompass.txt", "-base64", "120"])

  subprocess.run(["openssl", "enc", "-aes-256-cbc", "-salt", "-in", data, "-out",   new_file, "-pass", "file:randompass.txt"])


  subprocess.run(["openssl", "rsautl","-encrypt", "-in", "randompass.txt", "-out"  , "randompass.enc", "-pubin", "-inkey",  pubkey])

# Decryption

elif (answer == "Decrypt"):

  privkey = input("What is the path to      your private key? ")

  file_name = input("What the path of the encrypted data file? ")

  data = input("What would you like to name the new data file? ")

  subprocess.run(["openssl", "rsautl", "-decrypt", "-inkey", privkey, "-in", "randompass.enc","-out", "decodedpass.txt"])


  subprocess.run(["openssl","enc", "-d",    "-aes-256-cbc", "-in", file_name,"-out",  data, "-pass", "file:decodedpass.txt"])

else:
  print("Try Again. Input is case sensitive")
