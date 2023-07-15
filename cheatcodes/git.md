# Git

- WIP commit:
    ```
    git add .
    git commit -m "wip"
    git switch otherbranch
    ```
    And then later
    ```
    git switch firstbranch
    git reset --soft @~1
    ```
- SignOff:
  ```
  git commit -s -m "commit message"
  ```