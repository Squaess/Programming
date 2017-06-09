#ifndef INSERTWINDOW_H
#define INSERTWINDOW_H

#include <QMainWindow>

namespace Ui {
class InsertWindow;
}

class InsertWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit InsertWindow(QWidget *parent = 0);
    ~InsertWindow();

signals:
    void plsAddStud(QString data);

private slots:
    void on_addB_clicked();

private:
    Ui::InsertWindow *ui;

};

#endif // INSERTWINDOW_H
