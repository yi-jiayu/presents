from itertools import product

from pypresent import Present

keys = ('0' * 32, 'F' * 32)
plaintexts = ('0' * 16, 'F' * 16)

if __name__ == '__main__':
    cases = product(keys, plaintexts)
    for k, p in cases:
        key = k.decode('hex')
        plain = p.decode('hex')
        cipher = Present(key)
        encrypted = cipher.encrypt(plain)
        c = encrypted.encode('hex')
        print '''	{{
\t\tKey:        "{}",
\t\tPlaintext:  "{}",
\t\tCiphertext: "{}",
\t}},'''.format(k, p, c)
