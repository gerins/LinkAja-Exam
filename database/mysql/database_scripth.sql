-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema LinkAja
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema LinkAja
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `LinkAja` DEFAULT CHARACTER SET utf8 ;
USE `LinkAja` ;

-- -----------------------------------------------------
-- Table `LinkAja`.`customer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `LinkAja`.`customer` (
  `CUSTOMER_NUMBER` INT NOT NULL AUTO_INCREMENT,
  `NAME` VARCHAR(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`CUSTOMER_NUMBER`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `LinkAja`.`account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `LinkAja`.`account` (
  `ACCOUNT_NUMBER` INT NOT NULL AUTO_INCREMENT,
  `CUSTOMER_NUMBER` INT NOT NULL,
  `BALANCE` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`ACCOUNT_NUMBER`),
  INDEX `fk_account_customer_idx` (`CUSTOMER_NUMBER` ASC) VISIBLE,
  CONSTRAINT `fk_account_customer`
    FOREIGN KEY (`CUSTOMER_NUMBER`)
    REFERENCES `LinkAja`.`customer` (`CUSTOMER_NUMBER`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

ALTER TABLE account AUTO_INCREMENT=555001;
ALTER TABLE customer AUTO_INCREMENT=1001;

INSERT INTO customer(NAME) VALUE
("Bob Martin"), ("Linus Torvalds");

INSERT INTO account(CUSTOMER_NUMBER, BALANCE) VALUE
(1001, 10000), (1002, 15000);
