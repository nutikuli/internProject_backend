/*
  Warnings:

  - You are about to drop the column `entityId` on the `File` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE `File` DROP FOREIGN KEY `AccountFileRelation`;

-- DropForeignKey
ALTER TABLE `File` DROP FOREIGN KEY `BankFileRelation`;

-- DropForeignKey
ALTER TABLE `File` DROP FOREIGN KEY `OrderFileRelation`;

-- DropForeignKey
ALTER TABLE `File` DROP FOREIGN KEY `ProductFileRelation`;

-- AlterTable
ALTER TABLE `File` DROP COLUMN `entityId`,
    ADD COLUMN `bankId` INTEGER NULL,
    ADD COLUMN `productId` INTEGER NULL;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `File_accountId_fkey` FOREIGN KEY (`accountId`) REFERENCES `Account`(`id`) ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `File_productId_fkey` FOREIGN KEY (`productId`) REFERENCES `Product`(`id`) ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `File_bankId_fkey` FOREIGN KEY (`bankId`) REFERENCES `Bank`(`id`) ON DELETE SET NULL ON UPDATE CASCADE;
