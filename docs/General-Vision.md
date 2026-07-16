```
                                   +----------------+
                                   |      USER      |
                                   +----------------+
                                   | id             |
                                   | name           |
                                   | email          |
                                   | settings       |
                                   | createdAt      |
                                   +----------------+
                                           |
                             1 ------------ * 
                                           |
                                           v
                                   +----------------+
                                   |    ACCOUNT     |
                                   +----------------+
                                   | id             |
                                   | institution    |
                                   | name           |
                                   | type           |
                                   | currency       |
                                   | initialBalance |
                                   | currentBalance |
                                   | isActive       |
                                   +----------------+
                                           |
                 +-------------------------+--------------------------+
                 |                                                    |
                 |                                                    |
                 v                                                    v
      +----------------------+                          +----------------------+
      |  FINANCIAL EVENT     |                          |    TRANSACTION       |
      +----------------------+                          +----------------------+
      | id                   |                          | id                   |
      | accountId            |                          | eventId (optional)   |
      | categoryId           |                          | accountId            |
      | title                |                          | amount               |
      | description          |                          | date                 |
      | type                 |                          | balanceAfter         |
      | direction            |                          | status               |
      | amount               |                          | description          |
      | recurrence           |                          +----------------------+
      | startDate            |
      | endDate              |
      | status               |
      +----------------------+
                 |
                 |
      +----------+-----------+
      |                      |
      v                      v
+---------------+     +----------------+
|   CATEGORY    |     |      RULE      |
+---------------+     +----------------+
| id            |     | frequency      |
| name          |     | interval       |
| icon          |     | dayOfMonth     |
| color         |     | dayOfWeek      |
| type          |     | occurrences    |
+---------------+     | until          |
                      | skipWeekend    |
                      +----------------+

        USER
          |
          +---------------------+
          |                     |
          v                     v
 +----------------+     +----------------+
 |      GOAL      |     |    SCENARIO    |
 +----------------+     +----------------+
 | id             |     | id             |
 | name           |     | name           |
 | targetAmount   |     | description    |
 | deadline       |     | isDefault      |
 | priority       |     +----------------+
 | status         |              |
 +----------------+              |
                                 |
                                 |
                         +-------+-------+
                         |               |
                         v               v
                Financial Events   Forecast Engine
                        |
                        v
                 Projected Timeline
```