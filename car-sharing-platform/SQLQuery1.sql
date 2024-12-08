USE CarSharingDB;
GO
-- Create Trigger for Membership Tier Upgrade
CREATE OR ALTER TRIGGER trg_UpgradeMembershipTier
ON Users
AFTER UPDATE
AS
BEGIN
    SET NOCOUNT ON;

    UPDATE Users
    SET MembershipTierID = 
        (SELECT TOP 1 ID
         FROM MembershipTiers
         WHERE TotalSpending >= SpendingThreshold
         ORDER BY SpendingThreshold DESC)
    WHERE ID IN (SELECT DISTINCT ID FROM Inserted);
END;