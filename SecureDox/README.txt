Purpose:
The purpose of this scrypt is to aid in the administering a hybrid cryprosystem.

Encryption:
This script utilizes the openssl tool set to produce a random 120 byte passphrase that is then used to encrypt a large data file using AES 256 Bit encryption. The passphrase is then encrypted with the recipients public key using RSA encryption. This allows a secure method of transfer of  both the encrypted data file and encrypted passphrase to the recipient. The recipient may then decrypt the passphrase using his/her RSA private key, and then in turn decrypt the data file using the passphrase.

Decryption:
Upon recieving the encrypted data file and encrypted passphrase from the sender, the user can use this script to decode the passphrase with his/her RSA private key.The newly decrypted passphrase is then used to decrypt the data file allowing a secure method of data transfer.
