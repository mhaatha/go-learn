# CLI Task Manager

# Functional Requirements

### FR-01 - Menambahkan Task

Command:

```bash
tasker add "Learn Go for 2 hours straight"
```

Aplikasi harus:

1. Membuat task baru.
2. Memberikan ID unik kepada task.
3. Menyimpan waktu pembuatan task.
4. Menyimpan status awal sebagai `pending`.
5. Menyimpan task secara persistent.

Output Example:

```bash
Task created successfully.

ID: 1
Title: Learn Go for 2 hours straight
Status: pending
```

### FR-02 - Melihat Semua Task

Command:

```bash
tasker list
```

Aplikasi harus:

1. Menampilkan task dengan ID, status, dan title.
2. Task ditampilkan secara konsisten.
3. Jika belum ada task, tampilkan pesan yang sesuai.

Output Example:

```bash
ID  STATUS     TITLE
1   pending    Learn Go for 2 hours straight
2   pending    Build REST API
3   completed  Read Go book
```

Jika tidak ada task, tampilkan:

```bash
No tasks found.
```

### FR-03 - Menyelesaikan Task

Command:

```bash
tasker done 1
```

Aplikasi harus:

1. Menyelesaikan task dengan mengubah status menjadi `completed`.

Output Example:

```bash
Task 1 marked as completed.
```

Jika task dengan ID yang diberikan tidak ditemukan, tampilkan:

```bash
Error: task 1 not found.
```

Jika task dengan ID yang diberikan sudah selesai, tampilkan:

```bash
Error: task 1 is already completed.
```

### FR-04 - Menghapus Task

Command:

```bash
tasker delete 2
```

Aplikasi harus:

1. Meminta konfirmasi sebelum menghapus task.
2. Menghapus task dengan ID yang diberikan.

Output Example:

```bash
Delete task 2 "Build REST API"? [y/N]
```

Jika user memilih `y`, maka:

```bash
Task 2 deleted successfully.
```

Jika user memilih selain `y`, maka:

```bash
Task deleteion cancelled.
```

### FR-05 - Mengubah Task

Command:

```bash
tasker edit 1 "Learn Go Concurrency"
```

Aplikasi harus:

1. Mengubah task dengan ID yang diberikan.
2. Status task dan waktu pembuatannya tidak boleh berubah.

Output Example:

```bash
Task 1 updated successfully.
```

### FR-06 - Melihat Satu Task

Command:

```bash
tasker show 1
```

Aplikasi harus:

1. Menampilkan task dengan ID, status, title, dan waktu pembuatan.

Output Example:

```bash
ID:        1
Title:     Learn Go
Status:    pending
Created:   2026-09-16 20:15:32
```

Jika task dengan ID yang diberikan tidak ditemukan, tampilkan:

```bash
Error: task 1 not found.
```

# Command Specification

Command yang tersedia:

```bash
tasker add <title>
tasker list
tasker show <id>
tasker edit <id> <new-title>
tasker done <id>
tasker delete <id>
```

Jika pengguna menjalankan:

```bash
tasker
```

tanpa argument, aplikasi harus menampilkan bantuan penggunaan:

```bash
Tasker - CLI Task Manager

Usage:
  tasker <command> [arguments]

Commands:
  add       Create a new task
  list      List all tasks
  show      Show task details
  edit      Edit a task
  done      Mark a task as completed
  delete    Delete a task
  help      Show this help message
```

# Error Handling
Aplikasi harus menangani input yang tidak valid dengan baik.

Contoh:
```bash
tasker done
```
harus menghasilkan error yang jelas.
Contoh:
```bash
Error: missing task ID.
```
Contoh:
```bash
tasker done abc
```
```bash
Error: invalid task ID.
```

Contoh:
```bash
tasker unknown
```
```bash
Error: unknown command "unknown".
```

Aplikasi tidak boleh crash hanya karena pengguna memberikan input yang salah.
