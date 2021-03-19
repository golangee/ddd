# ManageTickets
As a SupportietyAdmin or Application 
I want to manage tickets
so that I can submit or delete new incidents.

## AdminDeletesTicket
As a SupportietyAdmin 
I want to delete tickets from a user identified by his SecId, 
so that I can comply to the DSGVO/GDPR.

### DoubleDelete
Given I'm a SupportietyAdmin
when I delete the same ticket twice
then I want a message telling me that its not possible.