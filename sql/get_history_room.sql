SELECT CONCAT(IFNULL(users.name_prefix,"") ,"",users.first_name," ",users.last_name) as fullname ,bookings.start_date ,bookings.end_date,rooms.room_name  
FROM users 
JOIN bookings ON users.id = bookings.users_id 
JOIN rooms ON bookings.rooms_id  = rooms.id 