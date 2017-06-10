#ifndef DELETEWINDOW_H
#define DELETEWINDOW_H

#include <QMainWindow>

namespace Ui {
class DeleteWindow;
}

class DeleteWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit DeleteWindow(QWidget *parent = 0);
    ~DeleteWindow();
signals:
    void plsDeleteStudent(QString data);
    void plsDeleteCourse(QString data);

private slots:
    void on_pushButton_clicked();
    void on_deleteB_clicked();

private:
    Ui::DeleteWindow *ui;
};

#endif // DELETEWINDOW_H
