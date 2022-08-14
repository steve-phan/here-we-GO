1. ### User _*struct*_

   - Name string
   - Age int

2. ### UserCollection _*struct*_

   - Users []\*User
   - _*implements*_ CreateIterator() **Iterator**

3. ### Interator _*interface*_

   - HasNext() _bool_
   - GetNext() _User_

4. ### UserIterator _*struct*_

   - index int
   - Users []\*User
   - _*implements*_
     - HasNext() _bool_
     - GetNext() _User_
