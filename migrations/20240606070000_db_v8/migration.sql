/*
  Warnings:

  - Added the required column `roleName` to the `AdminPermission` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `AdminPermission` ADD COLUMN `roleName` VARCHAR(191) NOT NULL;
