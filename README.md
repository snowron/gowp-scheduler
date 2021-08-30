### Make Your Whatsapp Messages Schedule

Package has two functional actions. _Instant or Schedule_

The whatsapp qr code will be on your terminal. So easily can scan and start use

# Instant Mode
``
go run . -type instant -contactsPath ./service/test-data/contacts.csv
``

The command runs the package with instant mode and next step is typing message.

````
Enter Message: Test message from cli
Test message from cli
Are You Sure? [y/N] :
````

Your csv file should be like mentioned below. Actually package doesn't use name field but sometimes saves us from typo or mismatch problems 
```
murat,90539xxxxxxx
turan,90539xxxxxxx
```

# Schedule Mode
``
go run . -type schedule -ordersPath ./service/test-data/orders.csv
``

Your csv file should be like mentioned below. 
```
90539,hello from csv,2022-10-13T12:12:00Z
90539,hello from go,2022-10-13T12:12:00Z
```