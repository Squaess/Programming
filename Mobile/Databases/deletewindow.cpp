#include "deletewindow.h"
#include "ui_deletewindow.h"

DeleteWindow::DeleteWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::DeleteWindow)
{
    ui->setupUi(this);
}

DeleteWindow::~DeleteWindow()
{
    delete ui;
}
//delete student
void DeleteWindow::on_pushButton_clicked()
{
    QString data = "";
    data = ui->sidF->text();
    emit plsDeleteStudent(data);

}

void DeleteWindow::on_deleteB_clicked()
{
    QString data = "";
    data = ui->cidF->text();
    emit plsDeleteCourse(data);
}
