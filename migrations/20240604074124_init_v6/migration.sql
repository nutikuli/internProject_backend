/*
  Warnings:

  - You are about to drop the `_StoreProductCategory` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `storeId` to the `ProductCategory` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE `_StoreProductCategory` DROP FOREIGN KEY `_StoreProductCategory_A_fkey`;

-- DropForeignKey
ALTER TABLE `_StoreProductCategory` DROP FOREIGN KEY `_StoreProductCategory_B_fkey`;

-- AlterTable
ALTER TABLE `ProductCategory` ADD COLUMN `storeId` INTEGER NOT NULL;

-- DropTable
DROP TABLE `_StoreProductCategory`;

-- AddForeignKey
ALTER TABLE `ProductCategory` ADD CONSTRAINT `ProductCategory_storeId_fkey` FOREIGN KEY (`storeId`) REFERENCES `Account`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;
