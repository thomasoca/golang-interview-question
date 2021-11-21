SELECT a.ID, a.UserName, b.UserName as ParentUserName
FROM User a
LEFT JOIN USER b
ON a.Parent = b.ID
ORDER BY 1
