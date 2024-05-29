/*
  Warnings:

  - You are about to drop the column `orderId` on the `Order` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX `Order_orderId_key` ON `Order`;

-- AlterTable
ALTER TABLE `Order` DROP COLUMN `orderId`;
