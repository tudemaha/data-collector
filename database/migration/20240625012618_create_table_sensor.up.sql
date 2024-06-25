CREATE TABLE sensors (
    `id` INT NOT NULL AUTO_INCREMENT,
    `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `humidity` FLOAT,
    `temp` FLOAT,
    `soil_moisture` FLOAT,
    `soil_ph` FLOAT,
    `gas` FLOAT,
    `latitude` VARCHAR(100),
    `langitude` VARCHAR(100),
    `image` VARCHAR(200),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;