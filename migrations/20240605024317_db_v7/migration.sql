/*
  Warnings:

  - You are about to drop the column `accountId` on the `Log` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE `Log` DROP FOREIGN KEY `Log_accountId_fkey`;

-- AlterTable
ALTER TABLE `Log` DROP COLUMN `accountId`;
