# Simple Database Querying

Given a USER table with 3 columns: ID, UserName, Parent; with ID is the primary key, UserName is the name of the user, and Parent is the ID of the User creator of a particular User, e.g.:

```
|ID|UserName|Parent|
--------------------
| 1 | Ali | 2 |
| 2 | Budi | 0 |
| 3 | Cecep | 1 |
```

Write a SQL query to obtain this particular data:

```
|ID|UserName|ParentUserName|
--------------------
| 1 | Ali | Budi |
| 2 | Budi | NULL |
| 3 | Cecep | Ali |
```

where ParentUserName column is the UserName based on the Parent value.

## Solution

We need to left join the USER table with itself to get the ParentUserName from the ID of the user and to show NULL values when the ID is not exists on the table, i.e.

```sql
SELECT a.ID, a.UserName, b.UserName as ParentUserName
FROM User a
LEFT JOIN USER b
ON a.Parent = b.ID
ORDER BY 1
```
