import sys
import os
from Crypto.Cipher import PKCS1_OAEP
from Crypto.PublicKey import RSA
from Crypto.Hash import SHA256


def generate_key(pem_file, label=None):

    key = RSA.importKey(open(pem_file, "r").read())

    return PKCS1_OAEP.new(key, hashAlgo=SHA256, label=label)


def encrypt(message, label=None):

    cipher = generate_key("public.pem")

    return cipher.encrypt(message)


def decrypt(message, label=None):

    cipher = generate_key("private.pem")

    return cipher.decrypt(message)


def main():

    message = ''.join(sys.argv[1:])
    encrypt_msg = encrypt(message)

    with open("encrypted_data.txt", "wb") as en_f:

        en_f.write(encrypt_msg)


    if os.path.exists("server_message.txt"):
       
        with open("server_message.txt", "rb") as de_f:

            message = de_f.read()
            print decrypt(message)

    else:
       
        print "No response from Server"


if __name__ == "__main__":

    main()
