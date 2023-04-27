SELECT CONCAT(IFNULL(users.name_prefix,"") ,"",users.first_name," ",users.last_name) as fullname ,bookings.start_date ,bookings.end_date,rooms.room_name  
FROM users 
JOIN bookings ON users.id = bookings.users_id 
JOIN rooms ON bookings.rooms_id  = rooms.id 
WHERE bookings.start_date BETWEEN  "2021-09-02 15:00:00" AND "2021-09-02 19:00:59"
OR bookings.end_date  BETWEEN  "2021-09-02 15:00:00" AND "2021-09-02 19:00:59"