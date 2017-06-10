#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QSqlDatabase>
#include <insertwindow.h>
#include <deletewindow.h>
#include <searchwindow.h>

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
    void on_ssB_clicked();

    void on_scB_clicked();

    void on_deleteB_clicked();

    void on_searchB_clicked();

public slots:
    void addStudent(QString data);
    void addGrade(QString data);
    void addCourse(QString data);
    void deleteStudent(QString data);
    void deleteCourse(QString data);
    void searchStudName(QString data);
    void searchStudSur(QString data);
    void searchCourse(QString data);

private:
    Ui::MainWindow *ui;
    QSqlDatabase db;
public:
    InsertWindow *iw;
    DeleteWindow *dw;
    SearchWindow *sw;
};

#endif // MAINWINDOW_H
