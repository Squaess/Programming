#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QSqlDatabase>
#include <insertwindow.h>

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();
public:
    bool openDB();
private slots:
    void on_showB_clicked();
    void on_insertB_clicked();
public slots:
    void addStudent(QString data);

private:
    Ui::MainWindow *ui;
    InsertWindow *iw;
    QSqlDatabase db;
};

#endif // MAINWINDOW_H
