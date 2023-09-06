CREATE USER 'coba'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'coba'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE `user` (
  `ID` int(11) NOT NULL,
  `username` varchar(32) NOT NULL,
  `passkey` varchar(64) NOT NULL,
  `balance` bigint(20) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`);

ALTER TABLE `user`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

INSERT INTO `user` (`ID`, `username`, `passkey`, `balance`) VALUES
(1, 'johndoe', '$2a$14$uQpPUQ5NgR7M1BEZeHiOcO0syd8oHPTPNTfDHcVtbwO5n4M57EWo6', 100000000);
-- passkey: secret